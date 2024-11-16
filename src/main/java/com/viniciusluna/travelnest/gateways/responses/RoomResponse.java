package com.viniciusluna.travelnest.gateways.responses;

import lombok.Data;

import java.time.LocalDateTime;
import java.util.Date;

@Data
public class RoomResponse  {
    private String id;
    private int number;
    private int floor;
    private String roomsCategories;
    private float pricePerNight;
    private boolean available;
    private LocalDateTime createdAt;
    private LocalDateTime updatedAt;
}