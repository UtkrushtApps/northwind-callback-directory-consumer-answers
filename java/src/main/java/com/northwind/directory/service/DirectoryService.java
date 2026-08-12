package com.northwind.directory.service;

import com.northwind.directory.client.CountriesClient;
import com.northwind.directory.client.CountryPayload;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class DirectoryService {

    private final CountriesClient client;

    public DirectoryService(CountriesClient client) {
        this.client = client;
    }

    public String getPhoneNumbers(String country, String phone) {
        CountryPayload payload = client.fetchCountry(country);
        if (payload == null || payload.getData() == null || payload.getData().isEmpty()) {
            return "-1";
        }

        CountryPayload.CountryRecord record = payload.getData().get(0);
        List<String> callingCodes = record.getCallingCodes();
        if (callingCodes == null || callingCodes.isEmpty()) {
            return "-1";
        }

        String callingCode = callingCodes.get(callingCodes.size() - 1);
        return "+" + callingCode + " " + phone;
    }
}
