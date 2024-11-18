package com.viniciusluna.travelnest.gateways.controllers;

import com.viniciusluna.travelnest.domain.User;
import com.viniciusluna.travelnest.gateways.repositories.UserRepository;
import com.viniciusluna.travelnest.gateways.requests.AuthRequest;
import com.viniciusluna.travelnest.gateways.responses.AuthResponse;
import com.viniciusluna.travelnest.usecases.impl.TokenServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/auth")
public class AuthController {
    @Autowired
    private AuthenticationManager authenticationManager;
    @Autowired
    private UserRepository repository;
    @Autowired
    private TokenServiceImpl tokenService;

    @PostMapping("/login")
    public ResponseEntity login(@RequestBody AuthRequest data){
        var usernamePassword = new UsernamePasswordAuthenticationToken(data.getEmail(), data.getPassword());
        var auth = this.authenticationManager.authenticate(usernamePassword);

        var token = tokenService.generateToken((User) auth.getPrincipal());

        return ResponseEntity.ok(new AuthResponse(token));
    }
}
