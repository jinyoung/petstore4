package petshop.domain;

import java.util.*;
import lombok.*;
import petshop.domain.*;
import petshop.infra.AbstractEvent;

@Data
@ToString
public class AvatarCreated extends AbstractEvent {

    private Long id;

    public AvatarCreated(Customer aggregate) {
        super(aggregate);
    }

    public AvatarCreated() {
        super();
    }
    // keep

}
