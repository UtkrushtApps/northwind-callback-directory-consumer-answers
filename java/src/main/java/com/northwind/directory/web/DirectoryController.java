package com.northwind.directory.web;

import com.northwind.directory.service.DirectoryService;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.Map;

@RestController
public class DirectoryController {

    private final DirectoryService service;

    public DirectoryController(DirectoryService service) {
        this.service = service;
    }

    @GetMapping("/phone-numbers")
    public ResponseEntity<Map<String, String>> phoneNumbers(
            @RequestParam(required = false) String country,
            @RequestParam(required = false) String phone) {
        if (isBlank(country) || isBlank(phone)) {
            return ResponseEntity
                    .status(HttpStatus.BAD_REQUEST)
                    .body(Map.of("error", "country and phone are required"));
        }

        return ResponseEntity.ok(Map.of("result", service.getPhoneNumbers(country, phone)));
    }

    private boolean isBlank(String value) {
        return value == null || value.trim().isEmpty();
    }
}
