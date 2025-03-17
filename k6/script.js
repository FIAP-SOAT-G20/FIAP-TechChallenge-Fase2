import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
  vus: 10,
  duration: '30s',
};

// export const options = {
//     // Key configurations for avg load test in this section
//     vus: 100,
//     duration: '40s',
//     thresholds: {
//       http_req_duration: ['p(95)<250'], // 95% of requests must complete below 250ms
//     },
//     stages: [
//       { duration: '5m', target: 100 }, // traffic ramp-up from 1 to 100 users over 5 minutes.
//       { duration: '30m', target: 100 }, // stay at 100 users for 30 minutes
//       { duration: '5m', target: 0 }, // ramp-down to 0 users
//     ],
//   };

export default function() {
  let baseUrl = 'http://localhost:8080/api/v1';
  let res = http.get(`${baseUrl}/health`);
  check(res, { "status is 200": (res) => res.status === 200 });
  sleep(1);
}
