package com.viniciusluna.travelnest.usecases.interfaces;

import com.viniciusluna.travelnest.domain.User;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

@Service
public interface TokenService {
    String generateToken(User user);

    String validateToken(String token);
}
