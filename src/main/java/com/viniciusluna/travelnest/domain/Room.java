package com.viniciusluna.travelnest.domain;

import com.viniciusluna.travelnest.domain.enums.RoomsCategories;
import jakarta.persistence.*;
import lombok.Data;

import java.util.UUID;


@Data
@Entity
@Table(name = "tbl_rooms")
public class Room {
    @Id
    @GeneratedValue(generator = "UUID")
    private UUID id;

    @Column(name = "room_number", unique = true, nullable = false)
    private int number;

    @Column(name = "room_floor", nullable = false)
    private int floor;

    @Enumerated(EnumType.STRING)
    @Column(name = "room_category", nullable = false)
    private RoomsCategories roomsCategories;

    @Column(name = "room_price_per_night", nullable = false)
    private float pricePerNight;

    @Column(name = "room_available", nullable = false)
    private boolean available;
}
