import mongoose from 'mongoose';
import { connect, StringCodec } from 'nats';
import { StudentModel } from './data/models/student.js';

await mongoose.connect('mongodb://localhost:27017/cqrs');

const nc = await connect({
  url: 'nats://localhost:4222',
  json: true,
});

const sub = nc.subscribe('student.register');

const sc = StringCodec();

(async () => {
  for await (const m of sub) {
    let student = sc.decode(m.data);
    student = JSON.parse(student);

    await StudentModel.create(student);
  }
})();
