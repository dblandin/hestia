resource "aws_iam_role" "hestia" {
  name = "hestia-${terraform.workspace}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "hestia" {
  name        = "hestia-${terraform.workspace}"
  path        = "/"
  description = "IAM policy for hestia in ${terraform.workspace} environment."

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "lambda:InvokeFunction"
      ],
      "Resource": "*"
    },
    {
        "Effect": "Allow",
        "Action": [
            "ssm:DescribeParameters",
            "ssm:GetParameter",
            "ssm:GetParameters"
        ],
        "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "hestia" {
  role       = "${aws_iam_role.hestia.name}"
  policy_arn = "${aws_iam_policy.hestia.arn}"
}

resource "aws_iam_role_policy_attachment" "attachment" {
  role       = "${aws_iam_role.hestia.name}"
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
