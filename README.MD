# REBRAIN DevOps Task - Dev04. Writing API documentation.

## About The Project 
Hi, this is a Terraform `main.tf` file with __local-exec__ provisioner calling ansible `playbook.yml`. 
This playbook copy and then start __nginx__ web-server [http://oleg-web.devops.rebrain.srwx.net](http://oleg-web.devops.rebrain.srwx.net).

Terraform `main.tf` file creates 1 VPS and export `playbook.yml` ansible inventory file and then **Ansible** playbook with __nginx__ role comes into play.
File `api.go` contains apidoc instructions for creating simpleAPI documentation by `apidoc` command. Then ansible copy it to previously created VPS.

Link to autocreated documentation: [http://oleg-web.devops.rebrain.srwx.net/doc](http://oleg-web.devops.rebrain.srwx.net/doc)

All infrastructure you want to use for that can be created by **Terraform** on [GCP](https://www.cloud.google.com/) provider.

_Full task description represent at [REBRAIN site](https://lk.rebrainme.com/devops/task/164)_


### Built With
* [Terraform](https://www.terraform.io/)
* [Ansible](https://www.ansible.com/)
* [apidoc](https://apidocjs.com/)

## Getting started

### Prerequisites
Install all software you need:
  ```sh
  sudo apt install terraform
  sudo apt install ansible
  ```
## Installation
1. Clone the repo
   ```sh
   git clone https://gitlab.rebrainme.com/devops_users_repos/3732/dev04.git
   ```
## Usage
```sh
terraform init
terraform plan
terraform apply
```
## License

Distributed under the MIT License.

## Contact

Oleg Satalkin - [@olegen](https://t.me/olegen) - Oleg.Satalkin@gmail.com

Project Link: [https://gitlab.rebrainme.com/devops_users_repos/3732/dev04](https://gitlab.rebrainme.com/devops_users_repos/3732/dev04)

## Acknowledgments
* [REBRAIN](https://rebrainme.com/)