package petshop.external;

import java.util.Date;
import lombok.Data;

@Data
public class Item {

    private Long id;
    private String name;
    private String type;
    private Money price;
    private PetId petId;
    // keep

}
