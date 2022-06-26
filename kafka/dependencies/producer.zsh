docker run -it --rm \
  --network dependencies_default \
  --name producer \
  confluentinc/cp-kafkacat \
  kafkacat \
  -b kafka:9092 -P -t pinga

