# Data

data "template_file" "user_data_leader" {
  template = "${file("user-data-leader")}"

  vars {
    tpl_config_base_url = "${var.infrakit_config_base_url}",
    tpl_infrakit_group_suffix = "${random_id.group_suffix.hex}",
    tpl_image = "${var.do_image}",
    tpl_do_name = "${var.do_name}",
    tpl_region = "${var.do_region}",
    tpl_size = "${var.do_size}",
    tpl_ssh_key = "${var.do_ssh_key}",
  }
}

data "template_file" "user_data_manager" {
  template = "${file("user-data")}"
  vars {
    tpl_config_base_url = "${var.infrakit_config_base_url}",
    tpl_manager_ip = "${digitalocean_droplet.m1.ipv4_address_private}",
  }
}

# Resources

resource "random_id" "group_suffix" {
  byte_length = 8
}

resource "digitalocean_droplet" "m1" {
  image = "${var.do_image}"
  name = "${var.do_name}-manager1"
  region = "${var.do_region}"
  size = "${var.do_size}"
  ssh_keys = [ "${digitalocean_ssh_key.default.id}" ]
  tags = [ "${digitalocean_tag.swarm_role.id}", "${digitalocean_tag.project.id}" ]
  user_data = "${data.template_file.user_data_leader.rendered}"
  private_networking = "true"
}
resource "digitalocean_droplet" "m2" {
  image = "${var.do_image}"
  name = "${var.do_name}-manager2"
  region = "${var.do_region}"
  size = "${var.do_size}"
  ssh_keys = [ "${digitalocean_ssh_key.default.id}" ]
  tags = [ "${digitalocean_tag.swarm_role.id}", "${digitalocean_tag.project.id}" ]
  user_data = "${data.template_file.user_data_manager.rendered}"
}
resource "digitalocean_droplet" "m3" {
  image = "${var.do_image}"
  name = "${var.do_name}-manager3"
  region = "${var.do_region}"
  size = "${var.do_size}"
  ssh_keys = [ "${digitalocean_ssh_key.default.id}" ]
  tags = [ "${digitalocean_tag.swarm_role.id}", "${digitalocean_tag.project.id}" ]
  user_data = "${data.template_file.user_data_manager.rendered}"
}

resource "digitalocean_ssh_key" "default" {
  name = "ssh key"
  public_key = "${file("${var.do_ssh_key}")}"
}

resource "digitalocean_tag" "swarm_role" {
  name = "manager"
}

resource "digitalocean_tag" "project" {
  name = "${var.do_name}"
}

# Outputs

output "public-ip" {
  value = "${digitalocean_droplet.m1.ipv4_address}"
}
