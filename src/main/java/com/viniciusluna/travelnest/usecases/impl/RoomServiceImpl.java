package com.viniciusluna.travelnest.usecases.impl;

import com.viniciusluna.travelnest.domain.Room;
import com.viniciusluna.travelnest.gateways.repositories.RoomRepository;
import com.viniciusluna.travelnest.gateways.responses.RoomResponse;
import com.viniciusluna.travelnest.usecases.interfaces.RoomService;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@RequiredArgsConstructor
public class RoomServiceImpl implements RoomService {
    @Autowired
    private RoomRepository roomRepository;

    @Override
    public List<RoomResponse> findAllRooms(String available) {
        List<Room> rooms;

        if (available.equals("true")) {
            rooms = roomRepository.findByAvailable(true);
        } else if (available.equals("false")) {
            rooms = roomRepository.findByAvailable(false);
        } else {
            rooms = roomRepository.findAll();
        }

        return rooms.stream().map(this::mapToRoomResponse).toList();
    }

    private RoomResponse mapToRoomResponse(Room room) {
        var roomResponse = new RoomResponse();
        roomResponse.setId(room.getId().toString());
        roomResponse.setNumber(room.getNumber());
        roomResponse.setFloor(room.getFloor());
        roomResponse.setRoomsCategories(room.getRoomsCategories().name());
        roomResponse.setPricePerNight(room.getPricePerNight());
        roomResponse.setAvailable(room.isAvailable());
        roomResponse.setCreatedAt(room.getCreatedAt());
        roomResponse.setUpdatedAt(room.getUpdatedAt());
        return roomResponse;
    }
}
