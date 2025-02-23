import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { AssetsModule } from './assets/assets.module';
import { MongooseModule } from '@nestjs/mongoose';
import { WalletsModule } from './wallets/wallets.module';
import { OrdersModule } from './orders/orders.module';

@Module({
  imports: [
    MongooseModule.forRoot(
      'mongodb://root:root@mongo:27017/nest?authSource=admin&directConnection=true',
    ),
    AssetsModule,
    WalletsModule,
    OrdersModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
//REST - recursos
