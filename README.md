# Container-compose

A tool between docker-compose and helm, designed to help you play with distributed-systems with ease on your laptop **FOR DEVELOPMENT**.

# Warning

This is only a proof-of-concept, with barely a few go lines to back it up

# why not docker-compose and helm?

There is several issues with such softwares:

## Helm

* too much tied to kubernetes
* too complex to bootstrap a simple configuration
* no order

## docker-compose

* no order
* no notion of groups and templates. For example, if you need 3 zookeepers, you need to manually write 3 services

# Core idea

## groups

You can define groups, like this:

```yaml
zookeeper:
  image: bla/zookeeper
  number: 3
  templates:
    - /opt/zookeeper/conf/zk.conf
  ports:
    - {{ 2181 + group_number }}:2181
  hostname: {{ group_name }}-{{ group_number }}
  volumes:
    - data:
        path: /home/data/{{ group_name }}{{ group_number }}:/home/zk/conf
  command: /opt/zk/bin/start-foreground.sh
  ready_when:
    listening_port: 2181
    log: "is ready"

# order is important, ZK needs to be up before startintg the other group
hbase-master:
# ...
  links:
    - zookeeper
```

## variables

container-compose let you write termplates of configuration, and appply them BEFORE starting your application. It is also heavily using templating like this:

```
tickTime=2000
dataDir={{.volume.data}}
clientPort=2181
initLimit=5
syncLimit=2
# the syntax is badly wrong, but it is just here to give you insights
{{range .Group}}
server.{{.Number}}={{.hostname}}:2888:3888
{{end}}
# will produce
server.1=zoo1:2888:3888
server.2=zoo2:2888:3888
server.3=zoo3:2888:3888
```
To do such a thing, we need a daemon inside the docker tempalte that will:

* apply template
* start the command

Moby-compose should expose things like:

* IP
* current group name and id
* other groups hostnames

As such, we could inject some env variables that will be used by such a software(let's call it "container-compose"), for example:

* CONTAINER_COMPOSE_GROUP_NAME: zookeeper
* CONTAINER_COMPOSE_GROUP_NUMBER: 0
* CONTAINER_COMPOSE_TEMPLATES: /opt/zookeeper/conf/zk.conf,
* CONTAINER_COMPOSE_GROUP_LIST: zookeeper,hbase-master
* CONTAINER_COMPOSE_ZOOKEEPER_LIST: zookeeper-1, # generated with the hostnames rules
* ...

You could also list all the options that you want, and use for each on it, like:
* CONTAINER_COMPOSE_CUSTOM_BLABLA1: bla # will translate to key BLABLA1: bla