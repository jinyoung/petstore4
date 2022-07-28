package petshop.external;

import org.springframework.stereotype.Service;

@Service
public class PetServiceImpl implements PetService {

    /**
     * Fallback
     */
    public Pet getPet(Long id) {
        Pet pet = new Pet();
        return pet;
    }
}
