# sync-device

keeps device and sync status

## CQRS
 Command and queries has been separated. Commands are coming from Kafka with event driven architecture, queries are using http with GET apis.