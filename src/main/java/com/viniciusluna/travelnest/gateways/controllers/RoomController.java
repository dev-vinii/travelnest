package com.viniciusluna.travelnest.gateways.controllers;

import com.viniciusluna.travelnest.gateways.responses.RoomResponse;
import com.viniciusluna.travelnest.usecases.interfaces.RoomService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api/rooms")
public class RoomController {
    @Autowired
    private RoomService roomService;

    @GetMapping
    public List<RoomResponse> getRooms(@RequestParam(required = false, defaultValue = "all") String available) {
        return roomService.findAllRooms(available);
    }
}
