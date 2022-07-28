package petshop.domain;

import java.util.*;
import lombok.*;
import petshop.domain.*;
import petshop.infra.AbstractEvent;

@Data
@ToString
public class PetRegistered extends AbstractEvent {

    private Long id;
    private String name;
    private Object appearance;
    private Object energy;
    private String address;
    // keep

}
