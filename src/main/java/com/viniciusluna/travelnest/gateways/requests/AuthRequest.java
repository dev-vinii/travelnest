package com.viniciusluna.travelnest.gateways.requests;

import lombok.Data;

@Data
public class AuthRequest {
    private String email;
    private String password;
}
