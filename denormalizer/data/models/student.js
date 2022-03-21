import mongoose from 'mongoose';

class Student extends mongoose.Schema {
  constructor() {
    super(
      {
        id: {
          type: Number,
          required: true,
          index: true,
        },
        name: {
          type: String,
          required: true,
        },
        email: {
          type: String,
          required: true,
        },
        ra: {
          type: String,
          required: true,
        },
        phone: {
          type: String,
          required: true,
        },
        cpf: {
          type: String,
          required: true,
        },
        age: {
          type: Number,
          required: true,
        },
      },
      {
        timestamps: {
          createdAt: 'created_at',
          updatedAt: 'updated_at',
        },
      }
    );
  }
}

export const StudentModel = mongoose.model('students', new Student());
