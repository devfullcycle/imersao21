import { NestFactory } from '@nestjs/core';
import { AppModule } from '../app.module';
import { ConfluentKafkaServer } from '../kafka/confluent-kafka-server';

async function bootstrap() {
  const app = await NestFactory.createMicroservice(AppModule, {
    strategy: new ConfluentKafkaServer({
      server: {
        'bootstrap.servers': 'host.docker.internal:9094',
      },
      consumer: {
        allowAutoTopicCreation: true,
        sessionTimeout: 10000,
        rebalanceTimeout: 10000,
      },
    }),
  });
  console.log('Starting Kafka microservice');
  await app.listen();
}
bootstrap();
