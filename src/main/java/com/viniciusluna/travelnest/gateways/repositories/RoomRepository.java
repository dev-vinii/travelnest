package com.viniciusluna.travelnest.gateways.repositories;

import com.viniciusluna.travelnest.domain.Room;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;
import java.util.UUID;

public interface RoomRepository extends JpaRepository<Room, UUID> {
    List<Room> findByAvailable(boolean available);
}
