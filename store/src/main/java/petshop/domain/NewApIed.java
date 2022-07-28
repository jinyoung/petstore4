package petshop.domain;

import java.util.*;
import lombok.*;
import petshop.domain.*;
import petshop.infra.AbstractEvent;

@Data
@ToString
public class NewApIed extends AbstractEvent {

    private Long id;

    public NewApIed(Customer aggregate) {
        super(aggregate);
    }

    public NewApIed() {
        super();
    }
    // keep

}
