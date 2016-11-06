package dyndns

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/gregarmer/r53dyndns/config"
	"github.com/gregarmer/r53dyndns/utils"
)

type Dyndns struct {
	Config *config.Config
}

func (dyndns *Dyndns) getRoute53Client() *route53.Route53 {
	creds := credentials.NewStaticCredentials(dyndns.Config.AwsAccessKey,
		dyndns.Config.AwsSecretKey, "")

	conf := aws.NewConfig().WithRegion("us-east-1").WithCredentials(creds)

	return route53.New(session.New(conf))
}

func (dyndns *Dyndns) GetHostedZonesCount() {
	log.Printf("hello world")

	svc := dyndns.getRoute53Client()

	var params *route53.GetHostedZoneCountInput
	resp, err := svc.GetHostedZoneCount(params)
	utils.CheckErr(err)

	log.Printf("Zones: %d", *resp.HostedZoneCount)
}

func (dyndns *Dyndns) GetHostedZone() {
	svc := dyndns.getRoute53Client()

	params := &route53.GetHostedZoneInput{
		Id: aws.String(dyndns.Config.ZoneId),
	}
	resp, err := svc.GetHostedZone(params)
	utils.CheckErr(err)

	log.Printf(resp.String())
}

func (dyndns *Dyndns) GetResourceRecordSet(record string) string {
	svc := dyndns.getRoute53Client()

	params := &route53.ListResourceRecordSetsInput{
		HostedZoneId:          aws.String(dyndns.Config.ZoneId),
		MaxItems:              aws.String("20"),
		StartRecordIdentifier: aws.String("ResourceRecordSetIdentifier"),
		StartRecordName:       aws.String("DNSName"),
		StartRecordType:       aws.String("A"),
	}

	resp, err := svc.ListResourceRecordSets(params)
	utils.CheckErr(err)

	var recordResp *route53.ResourceRecordSet
	for _, rr := range resp.ResourceRecordSets {
		if aws.StringValue(rr.Type) != "A" {
			continue
		}

		rrName := aws.StringValue(rr.Name)
		rrName = strings.TrimSuffix(rrName, ".")

		if record == rrName {
			recordResp = rr
			break
		}
	}

	ip := *recordResp.ResourceRecords[0].Value
	return fmt.Sprintf("%s", ip)
}

func (dyndns *Dyndns) UpsertDomain(domain string, ip string) {
	svc := dyndns.getRoute53Client()

	existingIp := dyndns.GetResourceRecordSet(domain)
	if existingIp == ip {
		log.Printf("existing IP is already set to %s", ip)
		return
	}

	log.Printf("updating %s with IP %s", domain, ip)

	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String(route53.ChangeActionUpsert),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String(domain),
						Type: aws.String(route53.RRTypeA),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String(ip),
							},
						},
						TTL: aws.Int64(60),
					},
				},
			},
			Comment: aws.String("Updated by r53dyndns"),
		},
		HostedZoneId: aws.String(dyndns.Config.ZoneId),
	}
	resp, err := svc.ChangeResourceRecordSets(params)
	utils.CheckErr(err)
	log.Println(resp)

	// force a log entry
	log.SetOutput(os.Stdout)
	log.Printf("external IP for %s changed from %s to %s", domain, existingIp, ip)
}
