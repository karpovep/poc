const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const path = require('path');

// Завантаження .proto файлу
const PROTO_PATH = './src/protos/main.proto';

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
});
const exampleProto = grpc.loadPackageDefinition(packageDefinition).example;

// Реалізація сервісу
function getMessages(call, callback) {
    const messages = [
        { name: 'Athlete1', timeFinish: 123 },
        { name: 'Athlete2', timeFinish: 456 },
        { name: 'Athlete3', timeFinish: 789 },
    ];
    callback(null, { messages });
}

// Налаштування сервера
function main() {
    const server = new grpc.Server();
    server.addService(exampleProto.Olympic.service, { getMessages: getMessages });
    server.bindAsync('0.0.0.0:8080', grpc.ServerCredentials.createInsecure(), () => {
        console.log('Server running at http://127.0.0.1:8080');
        server.start();
    });
}

main();