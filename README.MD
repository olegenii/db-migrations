# REBRAIN DevOps Task - Dev06. Database migrations.

## About The Project
This project use `postgres-14` as database.

File `api.go` - is a web server with [API](http://api.oleg-web.devops.rebrain.srwx.net).
Check API docs on [http://oleg-web.devops.rebrain.srwx.net/doc](http://oleg-web.devops.rebrain.srwx.net/doc). It contains apidoc instructions for creating simple API documentation by `apidoc` command. Then ansible copy it to previously created VPS.

There are `makefile` and 6 sql migrations files created by [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) library in folder `db/migration` with seq prefix and description in their names.

up migrations:
1. `000001_init_schema.up.sql` - create `items` table
2. `000002_init_schema.up.sql` - create `warehouses` table and add `warehouse_id` column as FK to `items` table
3. `000003_sample_data.up.sql` - add sample data to `warehouse` and `items` tables

down migrations:
1. `000001_init_schema.down.sql` - drop `items` table
2. `000002_init_schema.down.sql` - drop `warehouses` table and drop `warehouse_id` column
3. `000003_sample_data.down.sql` - truncate `warehouse` and `items` tables


For 1 step up migration use
```sh
make migratestepup 
```
For 1 step down migration use
```sh
make migratestepdown 
```
For migration up
```sh
make migrateup 
```
For migration down
```sh
make migratedown 
```
For atabase drop 
```sh
make migratedrop
```

Terraform `main.tf` file creates 1 VPS and export `playbook.yml` ansible inventory file and then **Ansible** playbook with __nginx__, __postgres__, __go__ roles come into play.

Link to autocreated documentation: [http://oleg-web.devops.rebrain.srwx.net/doc](http://oleg-web.devops.rebrain.srwx.net/doc)

All infrastructure you want to use for that can be created by **Terraform** on [GCP](https://www.cloud.google.com/) provider.

_Full task description represent at [REBRAIN site](https://lk.rebrainme.com/devops/task/166)_


### Built With
* [Terraform](https://www.terraform.io/)
* [Ansible](https://www.ansible.com/)
* [apidoc](https://apidocjs.com/)
* [gorm](https://gorm.io/)

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
   git checkout dev06
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