package dyndns

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/gregarmer/r53dyndns/config"
	"github.com/gregarmer/r53dyndns/utils"
	"log"
)

type Dyndns struct {
	Config *config.Config
}

func (dyndns *Dyndns) GetRoute53Client() *route53.Route53 {
	creds := credentials.NewStaticCredentials(dyndns.Config.AwsAccessKey,
		dyndns.Config.AwsSecretKey, "")

	conf := aws.NewConfig().WithRegion("us-east-1").WithCredentials(creds)

	return route53.New(session.New(conf))
}

func (dyndns *Dyndns) GetHostedZonesCount() {
	log.Printf("hello world")

	svc := dyndns.GetRoute53Client()

	var params *route53.GetHostedZoneCountInput
	resp, err := svc.GetHostedZoneCount(params)
	utils.CheckErr(err)

	log.Printf("Zones: %d", *resp.HostedZoneCount)
}

func (dyndns *Dyndns) GetHostedZone() {
	svc := dyndns.GetRoute53Client()

	params := &route53.GetHostedZoneInput{
		Id: aws.String(dyndns.Config.ZoneId),
	}
	resp, err := svc.GetHostedZone(params)
	utils.CheckErr(err)

	log.Printf(resp.String())
}

func (dyndns *Dyndns) UpsertDomain(domain string, ip string) {
	log.Printf("upserting %s with IP %s", domain, ip)

	svc := dyndns.GetRoute53Client()

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
						// SetIdentifier: aws.String("ResourceRecordSetIdentifier"),
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
}

// func (route53 *Route53) GetAuth() aws.Auth {
// 	// setup aws auth
// 	auth := aws.Auth{}
// 	auth.AccessKey = route53.Config.AwsAccessKey
// 	auth.SecretKey = route53.Config.AwsSecretKey
// 	return auth
// }
//
// func (route53 *Route53) getOrCreateDomain() *route53.Domain {
// 	auth := route53.GetAuth()
//
// 	s := s3.New(auth, aws.USEast)
// 	bucket := s.Bucket(awsS3.Config.S3Bucket)
//
// 	if !awsS3.bucketExists {
// 		exists, err := bucket.Exists("")
// 		utils.CheckErr(err)
// 		if !exists {
// 			if !*noop {
// 				log.Printf("creating bucket %s", awsS3.Config.S3Bucket)
// 				err := bucket.PutBucket(s3.BucketOwnerFull)
// 				utils.CheckErr(err)
// 			} else {
// 				log.Printf("would create bucket %s (noop)", awsS3.Config.S3Bucket)
// 			}
// 		}
// 		awsS3.bucketExists = true
// 	}
//
// 	return bucket
// }
