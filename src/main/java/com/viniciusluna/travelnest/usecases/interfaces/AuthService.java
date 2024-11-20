package com.viniciusluna.travelnest.usecases.interfaces;

import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

@Service
public interface AuthService {
    ResponseEntity login(String email, String password);
}
