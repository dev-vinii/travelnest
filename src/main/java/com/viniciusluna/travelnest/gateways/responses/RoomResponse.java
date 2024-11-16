package com.viniciusluna.travelnest.gateways.responses;

import lombok.Data;

@Data
public class RoomResponse  {
    private String id;
    private int number;
    private int floor;
    private String roomsCategories;
    private float pricePerNight;
    private boolean available;
}