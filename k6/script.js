import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 20 }, // Ramp-up para 20 usuários em 30 segundos
    { duration: '10m', target: 20 }, // Manter 20 usuários por 1 minuto
    { duration: '30s', target: 0 }, // Ramp-down para 0 usuários em 30 segundos 
  ],
  thresholds: {
    http_req_duration: ['p(95)<500'], // 95% das requisições devem ser completadas em menos de 500ms
    'http_req_duration{staticAsset:yes}': ['p(95)<100'], // 95% dos assets estáticos devem ser completados em menos de 100ms
    http_req_failed: ['rate<0.01'], // Menos de 1% das requisições podem falhar
  }
};

export default function() {
    const BASE_URL = 'https://quickpizza.grafana.com';
    const response = http.get(`${BASE_URL}`);
    check(response, {
        'status is 200': (r) => r.status === 200,
        'tempo de resposta < 200ms': (r) => r.timings.duration < 200
    });
    sleep(1);
}
