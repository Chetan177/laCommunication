# Ansible
### Ansible is a radically simple IT automation engine that automates cloud provisioning, configuration management, application deployment, intra-service orchestration, and many other IT needs.

### Install Ansible

```
# Debian / Ubuntu
    sudo apt install ansible
# CentOs / AWS ec2-linux
    yum install ansible
```

### Setup Host machines
```
vi /etc/ansible/hosts
# Remove <hostip>, <hostUser> and <sshKey> with appropriate values

    [freeswitch_machine]
    <hostip> ansible_ssh_user=<hostUser> ansible_ssh_key=<sshKey>
```

### Test ansible connection to host machine
```
ansible-playbook freeswitch_machine -m ping
```

### Run ansible playbook
```
ansible-playbook freeswitch.yml
```