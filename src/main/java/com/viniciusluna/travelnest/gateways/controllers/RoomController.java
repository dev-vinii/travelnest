package com.viniciusluna.travelnest.gateways.controllers;

import com.viniciusluna.travelnest.gateways.requests.RoomRequest;
import com.viniciusluna.travelnest.gateways.responses.RoomResponse;
import com.viniciusluna.travelnest.usecases.interfaces.RoomService;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RequiredArgsConstructor
@RestController
@RequestMapping("/api/rooms")
public class RoomController {
    private final RoomService roomService;

    @GetMapping
    public List<RoomResponse> getRooms(@RequestParam(required = false, defaultValue = "all") String available) {
        return roomService.findAllRooms(available);
    }

    @PostMapping
    public ResponseEntity<?> createRoom(@RequestBody @Valid RoomRequest room, BindingResult validationResult) {
        if (validationResult.hasErrors()) {
            List<String> errors = validationResult.getAllErrors().stream().map(error -> error.getDefaultMessage()).toList();
            return ResponseEntity.badRequest().body(errors);
        }

        RoomResponse response = roomService.createRoom(room);
        return ResponseEntity.ok(response);
    }

}
