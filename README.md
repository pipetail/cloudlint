# cloudlint
----

cloudlint is an open source system for checking AWS configuration. It provides basic visibility into cloud spend, performance and security.

cloudlint is built upon many years of experience at running production workloads on AWS at scale, combined with best practices from the community, wrapped in a an user-friendly CLI.

----

## How to start using cloudlint

See our documentation at [docs.cloudlint.io](http://docs.cloudlint.io).

To use cloudlint code as a library in other applications, use the `github.com/pipetail/cloudlint` module.

### Installation
```
docker pull pipetail/cloudlint:latest
```

### Quick start
**Disclaimer**: mount your `~/.aws` folder if you really know what you are doing and can bear the security risks!
```
docker run -it --rm             \
    -v ~/.aws:/root/.aws:ro     \
    -e AWS_PROFILE=myawsprofile \
    -e AWS_SDK_LOAD_CONFIG=1    \
    pipetail/cloudlint:latest run
```

you can expect something like this:
```
+-----------------------+-------------------------+-------------------------------------+----------+----------+
| #                     | GROUP                   | NAME                                | IMPACT $ | SEVERITY |
+-----------------------+-------------------------+-------------------------------------+----------+----------+
| aws_dms_unused        | Resources with no usage | AWS DMS Replication instances       |        0 |     INFO |
| total_ebs_unused      | Resources with no usage | EBS Unused                          |     1501 |    ERROR |
| aws_paid_support      | Resources with no usage | AWS Paid Support                    |      200 |     INFO |
| aws_ebs_snapshots_old | Resources with no usage | EBS snapshots are old               |      259 |  WARNING |
| total_elb_unused      | Resources with no usage | ELB Unused                          |      160 |    ERROR |
| vpc_eps_notused       | Incorrect service usage | VPC S3 endpoints are not used       |        0 |  WARNING |
| nat_gw_unused         | Incorrect service usage | NAT Gateways are unused             |      390 |    ERROR |
| eip_unused            | Incorrect service usage | Elastic IPs are unused              |      112 |    ERROR |
| ami_old               | Resources with no usage | AMIs are too old                    |      657 |  WARNING |
| ebs_opt               | Incorrect service usage | EC2 instances are not EBS Optimized |        0 |  WARNING |
+-----------------------+-------------------------+-------------------------------------+----------+----------+
|                       |                         | TOTAL IMPACT                        |     3279 |          |
|                       |                         | AWS MONTHLY BILL                    |    17862 |          |
+-----------------------+-------------------------+-------------------------------------+----------+----------+
```

Yay, you can save up to `$3279` and improve your performace and security!

## How to start developing cloudlint

If you want to build cloudlint right away:
```
mkdir -p $GOPATH/src/pipetail
cd $GOPATH/src/pipetail
git clone https://github.com/pipetail/cloudlint
cd cloudlint
make dep
make build
./bin/cloudlint
```

## Support

Please join the `#cloudlint` channel on the [CNCF Slack](http://slack.cncf.io/) to ask questions, discuss how cloudlint can fit into your workflow, or just chat.

If you have a feature request or found a bug,
please submit a Github Issue with all the details.
