package com.viniciusluna.travelnest.gateways.controllers;

import com.viniciusluna.travelnest.gateways.requests.AuthRequest;
import com.viniciusluna.travelnest.usecases.impl.AuthServiceImpl;
import com.viniciusluna.travelnest.usecases.impl.TokenServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/auth")
public class AuthController {
    @Autowired
    private AuthServiceImpl authService;

    @PostMapping("/login")
    public ResponseEntity login(@RequestBody AuthRequest body){
        return authService.login(body.getEmail(), body.getPassword());
    }
}
