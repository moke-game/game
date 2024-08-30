import {Client, StatusOK} from 'k6/net/grpc';
import {check, sleep} from 'k6';
import {makeParams} from "../common/common.js";

const client = new Client();
client.load(['./'], 'game-k6.proto');

const GRPC_ADDR = __ENV.SERVER_HOST || '127.0.0.1:8081';

export default function () {
    client.connect(GRPC_ADDR, {
        plaintext: true
    });
    const data = {
        uid: 'test',
        message: 'hello',
    };
    const params = makeParams("test");
    let response = client.invoke('game.pb.DemoService/Hi', data,params);
    console.log(response);
    check(response, {
        'status is OK': (r) => r && r.status === StatusOK,
    });
    client.close();
    sleep(1);
}


