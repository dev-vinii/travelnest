package com.viniciusluna.travelnest.usecases.interfaces;

import com.viniciusluna.travelnest.gateways.responses.RoomResponse;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public interface RoomService {
    List<RoomResponse> findAllRooms(String available);
}
