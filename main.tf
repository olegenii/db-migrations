terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "4.18.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
    null = {
      source = "hashicorp/null"
      version = "3.1.1"
    }
  }
}

# Configure GCP Provider
provider "google" {
  project = var.gcp_project_id
  region  = "europe-central2"
  zone    = "europe-central2-a"
  credentials = file("key.json")
}

# Configure AWS Provider
provider "aws" {
  region = "eu-central-1"
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
}

# Create a default firewall rule
resource "google_compute_firewall" "default" {
  name    = "default-firewall"
  network = google_compute_network.vps_network.id

  allow {
    protocol = "icmp"
  }

  allow {
    protocol = "tcp"
    ports    = ["80", "22", "8080", "5432"]
  }

  source_tags = ["web"]
  source_ranges = ["0.0.0.0/0"]
  direction = "INGRESS"
}

# Create an internal VPC for inctances
resource "google_compute_network" "vps_network" {
  name = "vpc-network"
}

# Create an instance group with webservers and named port
resource "google_compute_instance_group" "backend" {
  name        = "backend"
  description = "Backend webserver instance group"
  
  instances = [for vm in google_compute_instance.vm_instance : vm.id]

  named_port {
    name = "http"
    port = "80"
  }
}

# Create a VPS for backend webserver
resource "google_compute_instance" "vm_instance" {
  for_each = toset(var.vps_list)
  name         = "${each.key}"
  machine_type = "f1-micro"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  metadata = {
    ssh-keys = "${var.gcp_ssh_user}:${file(var.gcp_ssh_pub_key_file)}"
  }

  network_interface {
    network = google_compute_network.vps_network.id
    access_config {
    }
  }

  tags = ["web"]
}

# Get specified DNS zone
data "aws_route53_zone" "selected" {
  name = var.aws_route53_zone
}

# Create DNS record for backend webservers
resource "aws_route53_record" "web" {
  for_each = google_compute_instance.vm_instance
  zone_id = data.aws_route53_zone.selected.zone_id
  name = each.value.name
  type    = "A"
  ttl     = "300"
  records = [each.value.network_interface.0.access_config.0.nat_ip]
}

# Create DNS record for api backend webservers
resource "aws_route53_record" "api" {
  for_each = google_compute_instance.vm_instance
  zone_id = data.aws_route53_zone.selected.zone_id
  name = "api.${each.value.name}"
  type    = "A"
  ttl     = "300"
  records = [each.value.network_interface.0.access_config.0.nat_ip]
}

# Create an inventory using template
resource "local_file" "vps" {
  filename = "${path.module}/${var.file_out}"
  content  = templatefile("${path.module}/${var.file_in}", {domain = var.aws_route53_zone, vps_list = google_compute_instance.vm_instance, gcp_user = var.gcp_ssh_user})
}

# Create a null resource for ansible call
resource "null_resource" "vps_ready" {

  provisioner "local-exec" {
    command = "ansible-playbook -i inventory.yml playbook.yml"
  }
  # wait till inventory.yml get ready
  depends_on = [
    local_file.vps,
  ]
}
