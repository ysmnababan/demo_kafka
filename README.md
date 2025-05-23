# Kafka Demo

This project demonstrates a basic Kafka setup using Docker Compose and Go clients.

## Getting Started

1. **Start Kafka Broker**

   ```sh
   docker compose up -d
   ```

2. **Create a Topic**

   ```sh
   docker exec -it broker /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --create --topic helloworld --partitions 2 --replication-factor 1
   ```

3. **Run the Consumer**

   ```sh
   cd kafka-consumer
   go run main.go
   ```

4. **Run the Producer**

   ```sh
   cd kafka-producer
   go run main.go
   ```

## Useful Kafka Commands

- **Add partitions to a topic:**
  ```sh
  docker exec -it broker /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --alter --topic helloworld --partitions 3
  ```

- **Describe a topic:**
  ```sh
  docker exec -it broker /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --describe --topic helloworld
  ```

- **Delete a topic:**
  ```sh
  docker exec -it broker /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --delete --topic helloworld
  ```

> **Note:** You cannot reduce the number of partitions for a topic. To do so, delete and recreate the topic with the desired number of partitions.

## Reference

- [Confluent Go Client Quick Start](https://developer.confluent.io/get-started/go/#build-consumer)