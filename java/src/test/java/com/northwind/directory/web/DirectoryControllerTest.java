package com.northwind.directory.web;

import com.northwind.directory.client.CountriesClient;
import com.northwind.directory.client.CountryPayload;
import com.northwind.directory.service.DirectoryService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.context.TestConfiguration;
import org.springframework.context.annotation.Bean;
import org.springframework.test.web.servlet.MockMvc;

import java.util.List;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(DirectoryController.class)
class DirectoryControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @TestConfiguration
    static class StubConfig {
        @Bean
        CountriesClient countriesClient() {
            return country -> {
                if ("Puerto Rico".equals(country)) {
                    return payloadWith("Puerto Rico", List.of("1", "1787", "1939"));
                }
                if ("Atlantis".equals(country)) {
                    CountryPayload payload = new CountryPayload();
                    payload.setData(List.of());
                    return payload;
                }
                return payloadWith("Afghanistan", List.of("93"));
            };
        }

        private CountryPayload payloadWith(String name, List<String> codes) {
            CountryPayload.CountryRecord record = new CountryPayload.CountryRecord();
            record.setName(name);
            record.setCallingCodes(codes);
            CountryPayload payload = new CountryPayload();
            payload.setData(List.of(record));
            return payload;
        }

        @Bean
        DirectoryService directoryService(CountriesClient client) {
            return new DirectoryService(client);
        }
    }

    @Test
    void endpointReturnsJsonForKnownCountry() throws Exception {
        mockMvc.perform(get("/phone-numbers")
                        .param("country", "Afghanistan")
                        .param("phone", "656445445"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.result").value("+93 656445445"));
    }

    @Test
    void endpointReturnsJsonForMultiCodeCountry() throws Exception {
        mockMvc.perform(get("/phone-numbers")
                        .param("country", "Puerto Rico")
                        .param("phone", "123456789"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.result").value("+1939 123456789"));
    }

    @Test
    void endpointReturnsSuccessfulMinusOneForUnknownCountry() throws Exception {
        mockMvc.perform(get("/phone-numbers")
                        .param("country", "Atlantis")
                        .param("phone", "5551234"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.result").value("-1"));
    }

    @Test
    void endpointRejectsMissingCountry() throws Exception {
        mockMvc.perform(get("/phone-numbers")
                        .param("phone", "5551234"))
                .andExpect(status().isBadRequest())
                .andExpect(jsonPath("$.error").value("country and phone are required"));
    }

    @Test
    void endpointRejectsBlankPhone() throws Exception {
        mockMvc.perform(get("/phone-numbers")
                        .param("country", "Afghanistan")
                        .param("phone", "   "))
                .andExpect(status().isBadRequest())
                .andExpect(jsonPath("$.error").value("country and phone are required"));
    }
}
