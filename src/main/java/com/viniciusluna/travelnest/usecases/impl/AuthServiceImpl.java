package com.viniciusluna.travelnest.usecases.impl;

import com.viniciusluna.travelnest.domain.User;
import com.viniciusluna.travelnest.gateways.repositories.UserRepository;
import com.viniciusluna.travelnest.gateways.responses.AuthResponse;
import com.viniciusluna.travelnest.usecases.interfaces.AuthService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

@Service
public class AuthServiceImpl implements AuthService {
    @Autowired
    private UserRepository repository;
    @Autowired
    private PasswordEncoder passwordEncoder;
    @Autowired
    private TokenServiceImpl tokenService;

    public ResponseEntity login(String email, String password){
        User user = this.repository.findByEmail(email).orElseThrow(() -> new RuntimeException("User not found"));
        if(passwordEncoder.matches(password, user.getPassword())) {
            String token = this.tokenService.generateToken(user);
            return ResponseEntity.ok(new AuthResponse(token));
        }
        return ResponseEntity.badRequest().build();
    }
}
