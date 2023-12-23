package reportroutine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/aptos"
	"github.com/NetSepio/gateway/util/pkg/ipfs"
)

type ExtendedReport struct {
	models.Report
	UpVotes    int      `json:"upvotes"`
	DownVotes  int      `json:"downvotes"`
	NotSure    int      `json:"notSure"`
	TotalVotes int      `json:"totalVotes"`
	Status     string   `json:"status"`
	Tags       []string `json:"tags"`
	Images     []string `json:"images"`
}

func ProcessAndUploadReports() {
	db := dbconfig.GetDb()
	now := time.Now()

	var extendedReports []ExtendedReport
	if err := db.Model(&models.Report{}).
		Select(`reports.*, 
				array_agg(distinct report_tags.tag) filter (where report_tags.tag is not null) as tags,
				array_agg(distinct report_images.image_url) filter (where report_images.image_url is not null) as images,
                (SELECT COUNT(DISTINCT voter_id) FROM report_votes WHERE report_id = reports.id and vote_type = 'upvote') as up_votes,
                (SELECT COUNT(DISTINCT voter_id) FROM report_votes WHERE report_id = reports.id and vote_type = 'downvote') as down_votes,
                (SELECT COUNT(DISTINCT voter_id) FROM report_votes WHERE report_id = reports.id and vote_type = 'notsure') as not_sure,
                (SELECT COUNT(DISTINCT voter_id) FROM report_votes WHERE report_id = reports.id) as total_votes`).
		Where("end_meta_data_hash IS NULL AND end_time < ?", now).
		Joins("left join report_tags on report_tags.report_id = reports.id").
		Joins("left join report_images on report_images.report_id = reports.id").
		Find(&extendedReports).Error; err != nil {
		fmt.Printf("Error retrieving reports: %v\n", err)
		return
	}

	for _, report := range extendedReports {
		if now.After(report.EndTime) {
			if float64(report.UpVotes)/float64(report.TotalVotes) >= 0.51 {
				report.Status = "accepted"
			} else {
				report.Status = "rejected"
			}
		} else {
			report.Status = "running"
		}
		// Convert report details to JSON
		reportJSON, err := json.Marshal(report)
		if err != nil {
			fmt.Printf("Error marshaling report: %v\n", err)
			continue
		}

		// Upload JSON to IPFS
		res, err := ipfs.UploadToIpfs(bytes.NewReader(reportJSON), "report.json")
		if err != nil {
			fmt.Printf("Error uploading report to IPFS: %v\n", err)
			continue
		}

		// Call ResolveProposal with the metadata hashes
		txResultData, err := aptos.ResolveProposal(*report.MetaDataHash, res.Value.Cid)
		if err != nil {
			fmt.Printf("Error resolving proposal: %v\n", err)
			continue
		}

		// Update report with new end metadata hash and transaction data
		report.EndMetaDataHash = &res.Value.Cid
		report.EndTransactionHash = &txResultData.Result.TransactionHash
		report.EndTransactionVersion = &txResultData.Result.Version
		if err := db.Save(&report).Error; err != nil {
			fmt.Printf("Error updating report in the database: %v\n", err)
			continue
		}

		fmt.Printf("Report updated successfully with end metadata hash: %s\n", *report.EndMetaDataHash)
	}

	if len(extendedReports) == 0 {
		fmt.Println("No unfinished reports found.")
	}
}

func StartProcessingReportsPeriodically() {
	for {
		ProcessAndUploadReports()    // Run the processing function
		time.Sleep(10 * time.Second) // Wait for 10 seconds before the next run
	}
}
