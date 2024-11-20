package com.viniciusluna.travelnest.gateways.requests;

import com.viniciusluna.travelnest.domain.enums.RoomsCategories;
import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Positive;
import lombok.Data;

@Data
public class RoomRequest {
    private int number;
    private int floor;
    private RoomsCategories roomsCategories;

    @NotNull(message = "O preço por noite é obrigatório")
    @Positive(message = "O preço por noite precisa ser maior que zero")
    private Float pricePerNight;
    private boolean available;
}
