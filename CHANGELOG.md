# 1.0.0 (2025-04-06)


### Bug Fixes

* add latest tag to Docker image during build process ([87f81e1](https://github.com/dudubernardino/gobank/commit/87f81e12b3866147048e3606d4716f51463aa505))
* correct Docker tag command and streamline image output in CI/CD workflow ([bb739b7](https://github.com/dudubernardino/gobank/commit/bb739b71289e432f80f265b21e3ee8bffb0ee962))
* set infrastructure project directory and update environment variables for Terraform apply step ([2a79772](https://github.com/dudubernardino/gobank/commit/2a7977215d3934fc8e99c9e100b686234cce6005))
* streamline Terraform commands by setting working directory for init, plan, and apply steps ([8a0d82f](https://github.com/dudubernardino/gobank/commit/8a0d82f3b748d861d2c7ec50ed921b5a93049e18))
* update AWS credentials audience to correct endpoint ([b13f5e7](https://github.com/dudubernardino/gobank/commit/b13f5e7a440e189b1e6f80813da7e3e5f262c4cf))
* update CI/CD workflow to include latest Docker image tagging ([627c946](https://github.com/dudubernardino/gobank/commit/627c946df31d7d101f725c85e8412dbbe31ae48a))
* update CI/CD workflow to output Docker image for deployment ([210e236](https://github.com/dudubernardino/gobank/commit/210e236bec284427470fd733a5cda6a73991928c))
* update Docker image naming and port configuration in CI/CD and infrastructure files ([142dbc2](https://github.com/dudubernardino/gobank/commit/142dbc26e0c03939fc527344ab034864b0b53075))
* update GitHub token secret reference in CI/CD workflow ([1f671f6](https://github.com/dudubernardino/gobank/commit/1f671f67bf6ed42101f95e2dbc1ea4595f592bf4))
* update permissions to allow write access for contents in CI/CD workflow ([8a314c7](https://github.com/dudubernardino/gobank/commit/8a314c753f077459118e6f42a3d3300e24ed4d5b))
* update release job to depend on deploy job and remove redundant dependency installation ([724134c](https://github.com/dudubernardino/gobank/commit/724134c562f419219cb8542d64d65178e8211055))


### Features

* add debugging step for ECR image variable and expand IAM role policy actions ([dc11a84](https://github.com/dudubernardino/gobank/commit/dc11a84b24d0b76746eb0a4efc2ccc3be274b8dd))
* add image tag mutability and scan on push settings for ECR module ([76bb47d](https://github.com/dudubernardino/gobank/commit/76bb47d1f87d435734fa50ae13054a3281ab9cf5))
* add VPC permissions to Terraform IAM role policy ([5f8b6b0](https://github.com/dudubernardino/gobank/commit/5f8b6b074cb72d6eac7cd0c184d7d8eea3a92264))
* implement VPC, RDS, and App Runner infrastructure with security groups and IAM roles ([930eb5a](https://github.com/dudubernardino/gobank/commit/930eb5aa2334d47a962ef2daa0aa4e685c147b9c))
