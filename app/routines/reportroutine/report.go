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
	"github.com/lib/pq"
)

type ExtendedReport struct {
	models.Report
	UpVotesCal    int            `json:"upvotes"`
	DownVotesCal  int            `json:"downvotes"`
	NotSureCal    int            `json:"notSure"`
	TotalVotesCal int            `json:"totalVotes"`
	Tags          pq.StringArray `json:"tags" gorm:"type:text[]"`
	Images        pq.StringArray `json:"images" gorm:"type:text[]"`
}

func ProcessAndUploadReports() {
	db := dbconfig.GetDb()
	now := time.Now()

	var extendedReports []ExtendedReport
	if err := db.Model(&models.Report{}).
		Select(`reports.*, 
            array_agg(distinct report_tags.tag) filter (where report_tags.tag is not null) as tags,
            array_agg(distinct report_images.image_url) filter (where report_images.image_url is not null) as images,
            (SELECT COUNT(DISTINCT voter_id) FROM report_votes WHERE report_id = reports.id and vote_type = 'upvote') as up_votes_cal,
            (SELECT COUNT(DISTINCT voter_id) FROM report_votes WHERE report_id = reports.id and vote_type = 'downvote') as down_votes_cal,
            (SELECT COUNT(DISTINCT voter_id) FROM report_votes WHERE report_id = reports.id and vote_type = 'notsure') as not_sure_cal,
            (SELECT COUNT(DISTINCT voter_id) FROM report_votes WHERE report_id = reports.id) as total_votes_cal`).
		Joins("left join report_tags on report_tags.report_id = reports.id").
		Joins("left join report_images on report_images.report_id = reports.id").
		Where("end_meta_data_hash IS NULL AND end_time < ?", now).
		Group("reports.id").
		Scan(&extendedReports).Error; err != nil {
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
		if report.MetaDataHash == nil {
			fmt.Printf("metadatahash is nil: %v\n", report)
			continue
		}
		// Call ResolveProposal with the metadata hashes
		txResultData, err := aptos.ResolveProposal(*report.MetaDataHash, res.Value.Cid)
		if err != nil {
			fmt.Printf("Error resolving proposal: %v\n", err)
			continue
		}

		fmt.Printf("Transaction hash: %s\n", txResultData.Result.TransactionHash)

		if err := db.Model(&models.Report{}).Where("id = ?", report.ID).Updates(models.Report{
			UpVotes:               report.UpVotesCal,
			DownVotes:             report.DownVotesCal,
			NotSure:               report.NotSureCal,
			TotalVotes:            report.TotalVotesCal,
			Status:                report.Status,
			EndTransactionHash:    &txResultData.Result.TransactionHash,
			EndMetaDataHash:       &res.Value.Cid,
			EndTransactionVersion: &txResultData.Result.Version,
		}).Error; err != nil {
			fmt.Printf("Error updating report in the database: %v\n", err)
			continue
		}

	}

}

func StartProcessingReportsPeriodically() {
	for {
		ProcessAndUploadReports()    // Run the processing function
		time.Sleep(10 * time.Second) // Wait for 10 seconds before the next run
	}
}
