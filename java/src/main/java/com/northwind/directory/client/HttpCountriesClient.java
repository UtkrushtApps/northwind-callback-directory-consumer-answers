package com.northwind.directory.client;

import com.northwind.directory.config.UpstreamConfig;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;
import org.springframework.web.util.UriComponentsBuilder;

import java.net.URI;
import java.time.Duration;

@Component
public class HttpCountriesClient implements CountriesClient {

    private final RestTemplate restTemplate;
    private final UpstreamConfig config;

    public HttpCountriesClient(UpstreamConfig config, RestTemplateBuilder restTemplateBuilder) {
        this.config = config;
        this.restTemplate = restTemplateBuilder
                .setConnectTimeout(Duration.ofMillis(config.getTimeoutMs()))
                .setReadTimeout(Duration.ofMillis(config.getTimeoutMs()))
                .build();
    }

    @Override
    public CountryPayload fetchCountry(String country) {
        URI uri = UriComponentsBuilder
                .fromUriString(config.getBaseUrl())
                .queryParam("name", country)
                .build()
                .encode()
                .toUri();
        return restTemplate.getForObject(uri, CountryPayload.class);
    }
}
