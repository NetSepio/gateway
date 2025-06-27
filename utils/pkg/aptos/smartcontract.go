package aptos

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/NetSepio/gateway/utils/load"
)

func argS(s string) string {
	return "string:" + s
}

func argA(s string) string {
	return "address:" + s
}

type DelegateReviewParams struct {
	Voter         string
	MetaDataUri   string
	Category      string
	DomainAddress string
	SiteUrl       string
	SiteType      string
	SiteTag       string
	SiteSafety    string
}

var ErrMetadataDuplicated = errors.New("metadata already exist")

func DelegateReview(p DelegateReviewParams) (*TxResult, error) {
	command := fmt.Sprintf("move run --function-id %s::reviews::delegate_submit_review --max-gas %d --gas-unit-price %d --args", load.Cfg.APTOS_FUNCTION_ID, load.Cfg.GAS_UNITS, load.Cfg.GAS_PRICE)
	args := append(strings.Split(command, " "),
		argA(p.Voter), argS(p.MetaDataUri), argS(p.Category), argS(p.DomainAddress), argS(p.SiteUrl), argS(p.SiteType), argS(p.SiteTag), argS(p.SiteSafety), argS(""))
	cmd := exec.Command("aptos", args...)
	fmt.Println(strings.Join(args, " "))

	o, err := cmd.Output()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			if strings.Contains(string(o), "ERROR_METADATA_DUPLICATED(0x3)") {
				return nil, fmt.Errorf("%w: %w", ErrMetadataDuplicated, err)
			}
			return nil, fmt.Errorf("stderr: %s out: %s err: %w", err.Stderr, o, err)
		}
		return nil, fmt.Errorf("out: %s err: %w", o, err)
	}

	txResult, err := UnmarshalTxResult(o)
	return &txResult, err
}

var ErrMetadataNotFound = errors.New("metadata not found")

func DeleteReview(metaDataUri string) (*TxResult, error) {

	command := fmt.Sprintf("move run --function-id %s::reviews::delete_review --max-gas %d --gas-unit-price %d --args", load.Cfg.APTOS_FUNCTION_ID, load.Cfg.GAS_UNITS, load.Cfg.GAS_PRICE)
	args := append(strings.Split(command, " "), argS(metaDataUri))
	cmd := exec.Command("aptos", args...)
	fmt.Println(strings.Join(args, " "))

	o, err := cmd.Output()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			if strings.Contains(string(o), "EKEY_NOT_FOUND(0x10002)") {
				return nil, fmt.Errorf("%w: %w", ErrMetadataNotFound, err)
			}
			return nil, fmt.Errorf("stderr: %s out: %s err: %w", err.Stderr, o, err)
		}
		return nil, fmt.Errorf("out: %s err: %w", o, err)
	}

	txResult, err := UnmarshalTxResult(o)
	return &txResult, err
}

func DelegateMintNft(minter string) (*TxResult, error) {
	command := fmt.Sprintf("move run --function-id %s::vpnv2::delegate_mint_NFT --max-gas %d --gas-unit-price %d --args", load.Cfg.APTOS_FUNCTION_ID, load.Cfg.GAS_UNITS, load.Cfg.GAS_PRICE)
	args := append(strings.Split(command, " "), argA(minter))
	cmd := exec.Command("aptos", args...)
	fmt.Println(strings.Join(args, " "))

	o, err := cmd.Output()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("stderr: %s out: %s err: %w", err.Stderr, o, err)
		}
		return nil, fmt.Errorf("out: %s err: %w", o, err)
	}

	txResult, err := UnmarshalTxResult(o)
	return &txResult, err
}

func UploadArchive(siteUrl string, siteIpfsHash string) (*TxResult, error) {
	command := fmt.Sprintf("move run --function-id %s::reviews::archive_link --max-gas %d --gas-unit-price %d --args", load.Cfg.APTOS_FUNCTION_ID, load.Cfg.GAS_UNITS, load.Cfg.GAS_PRICE)
	args := append(strings.Split(command, " "), argS(siteUrl), argS(siteIpfsHash))
	cmd := exec.Command("aptos", args...)
	fmt.Println(strings.Join(args, " "))

	o, err := cmd.Output()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("stderr: %s out: %s err: %w", err.Stderr, o, err)
		}
		return nil, fmt.Errorf("out: %s err: %w", o, err)
	}

	txResult, err := UnmarshalTxResult(o)
	return &txResult, err
}

func SubmitProposal(proposer string, metadata string) (*TxResult, error) {
	command := fmt.Sprintf("move run --function-id %s::report_dao_v1::submit_proposal --max-gas %d --gas-unit-price %d --args", load.Cfg.APTOS_REPORT_FUNCTION_ID, load.Cfg.GAS_UNITS, load.Cfg.GAS_PRICE)
	args := append(strings.Split(command, " "), argA(proposer), argS(metadata))
	cmd := exec.Command("aptos", args...)
	fmt.Println(strings.Join(args, " "))

	o, err := cmd.Output()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("stderr: %s out: %s err: %w", err.Stderr, o, err)
		}
		return nil, fmt.Errorf("out: %s err: %w", o, err)
	}

	txResult, err := UnmarshalTxResult(o)
	return &txResult, err
}

func ResolveProposal(old_metadata string, metadata string) (*TxResult, error) {
	command := fmt.Sprintf("move run --function-id %s::report_dao_v1::resolve_proposal --max-gas %d --gas-unit-price %d --args", load.Cfg.APTOS_REPORT_FUNCTION_ID, load.Cfg.GAS_UNITS, load.Cfg.GAS_PRICE)
	args := append(strings.Split(command, " "), argS(old_metadata), argS(metadata))
	cmd := exec.Command("aptos", args...)
	fmt.Println(strings.Join(args, " "))

	o, err := cmd.Output()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("stderr: %s out: %s err: %w", err.Stderr, o, err)
		}
		return nil, fmt.Errorf("out: %s err: %w", o, err)
	}

	txResult, err := UnmarshalTxResult(o)
	return &txResult, err
}

func DeleteProposal(metadata string) (*TxResult, error) {
	command := fmt.Sprintf("move run --function-id %s::report_dao_v1::delete_proposal --max-gas %d --gas-unit-price %d --args", load.Cfg.APTOS_REPORT_FUNCTION_ID, load.Cfg.GAS_UNITS, load.Cfg.GAS_PRICE)
	args := append(strings.Split(command, " "), argS(metadata))
	cmd := exec.Command("aptos", args...)
	fmt.Println(strings.Join(args, " "))

	o, err := cmd.Output()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("stderr: %s out: %s err: %w", err.Stderr, o, err)
		}
		return nil, fmt.Errorf("out: %s err: %w", o, err)
	}

	txResult, err := UnmarshalTxResult(o)
	return &txResult, err
}
