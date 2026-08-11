import {Client, StatusOK} from 'k6/net/grpc';
import {check, sleep} from 'k6';
import {makeParams} from "../common/common.js";

const client = new Client();
client.load(['./'], 'game0-k6.proto', 'auth-k6.proto');

// Aggregate: AUTH_HOST can equal SERVER_HOST (in-process AuthService).
// Thin: set AUTH_HOST to the remote AuthService (e.g. 127.0.0.1:8082).
const GRPC_ADDR = __ENV.SERVER_HOST || '127.0.0.1:8081';
const AUTH_ADDR = __ENV.AUTH_HOST || GRPC_ADDR;
const AUTH_ID = __ENV.AUTH_ID || 'k6-user';
const APP_ID = __ENV.APP_ID || 'test';

export function setup() {
    client.connect(AUTH_ADDR, {plaintext: true});
    const authResp = client.invoke('auth.v1.AuthService/Authenticate', {
        id: AUTH_ID,
        app_id: APP_ID,
    });
    check(authResp, {
        'authenticate ok': (r) => r && r.status === StatusOK,
    });
    const token = authResp.message && (authResp.message.accessToken || authResp.message.access_token);
    client.close();
    if (!token) {
        throw new Error('Authenticate did not return access_token (check AUTH_HOST for thin topology)');
    }
    return {token};
}

export default function (data) {
    client.connect(GRPC_ADDR, {
        plaintext: true
    });
    const payload = {
        message: 'hello',
        topic: 'game',
    };
    const params = makeParams(data.token);
    const response = client.invoke('game0.pb.Game0Service/Hi', payload, params);
    console.log(response);
    check(response, {
        'status is OK': (r) => r && r.status === StatusOK,
    });
    client.close();
    sleep(1);
}
