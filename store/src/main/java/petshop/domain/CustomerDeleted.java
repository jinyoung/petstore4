package petshop.domain;

import java.util.*;
import lombok.*;
import petshop.domain.*;
import petshop.infra.AbstractEvent;

@Data
@ToString
public class CustomerDeleted extends AbstractEvent {

    private Long id;

    public CustomerDeleted(Customer aggregate) {
        super(aggregate);
    }

    public CustomerDeleted() {
        super();
    }
    // keep

}
