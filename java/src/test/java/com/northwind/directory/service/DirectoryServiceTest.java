package com.northwind.directory.service;

import com.northwind.directory.client.CountriesClient;
import com.northwind.directory.client.CountryPayload;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;

class DirectoryServiceTest {

    private CountriesClient stubClient(CountryPayload payload) {
        return country -> payload;
    }

    private CountryPayload payloadWith(String name, List<String> codes) {
        CountryPayload.CountryRecord record = new CountryPayload.CountryRecord();
        record.setName(name);
        record.setCallingCodes(codes);
        CountryPayload payload = new CountryPayload();
        payload.setData(List.of(record));
        return payload;
    }

    private CountryPayload emptyPayload() {
        CountryPayload payload = new CountryPayload();
        payload.setData(List.of());
        return payload;
    }

    @Test
    void resolvesSingleCallingCode() {
        DirectoryService service = new DirectoryService(
                stubClient(payloadWith("Afghanistan", List.of("93"))));
        assertEquals("+93 656445445", service.getPhoneNumbers("Afghanistan", "656445445"));
    }

    @Test
    void resolvesMultipleCallingCodesUsingHighestIndexCode() {
        DirectoryService service = new DirectoryService(
                stubClient(payloadWith("Puerto Rico", List.of("1", "1787", "1939"))));
        assertEquals("+1939 123456789", service.getPhoneNumbers("Puerto Rico", "123456789"));
    }

    @Test
    void returnsMinusOneForUnknownCountry() {
        DirectoryService service = new DirectoryService(stubClient(emptyPayload()));
        assertEquals("-1", service.getPhoneNumbers("Atlantis", "5551234"));
    }
}
