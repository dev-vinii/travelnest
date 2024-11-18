package com.viniciusluna.travelnest.usecases.interfaces;

import com.viniciusluna.travelnest.domain.User;
import org.springframework.stereotype.Service;

import java.time.Instant;

@Service
public interface TokenService {
    String generateToken(User user);

    String validateToken(String token);
}
