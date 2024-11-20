package com.viniciusluna.travelnest.domain;

import com.viniciusluna.travelnest.domain.enums.RoomsCategories;
import jakarta.persistence.*;
import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Positive;
import lombok.Data;
import org.hibernate.annotations.CreationTimestamp;
import org.hibernate.annotations.UpdateTimestamp;

import java.time.LocalDateTime;
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


    @NotNull(message = "O preço por noite é obrigatório")
    @Positive(message = "O preço por noite precisa ser maior que zero")
    @Column(name = "room_price_per_night", nullable = false)
    private Float pricePerNight;
    @Column(name = "room_available", nullable = false)
    private boolean available;

    @CreationTimestamp
    @Column(name = "room_created_at", nullable = false)
    private LocalDateTime createdAt;

    @UpdateTimestamp
    @Column(name = "room_updated_at", nullable = false)
    private LocalDateTime updatedAt;
}
