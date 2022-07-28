package petshop.domain;

import java.util.Date;
import java.util.List;
import javax.persistence.*;
import org.springframework.beans.BeanUtils;
import petshop.PetApplication;
import petshop.domain.PetRegistered;

@Entity
@Table(name = "Pet_table")
public class Pet {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;

    private String name;

    private int appearance;

    private int energy;

    @Embedded
    private Address address;

    @PostPersist
    public void onPostPersist() {
        PetRegistered petRegistered = new PetRegistered();
        BeanUtils.copyProperties(this, petRegistered);
        petRegistered.publishAfterCommit();
        // Get request from Item
        //petshop.external.Item item =
        //    Application.applicationContext.getBean(petshop.external.ItemService.class)
        //    .getItem(/** mapping value needed */);

    }

    @PrePersist
    public void onPrePersist() {
        // Get request from Item
        //petshop.external.Item item =
        //    Application.applicationContext.getBean(petshop.external.ItemService.class)
        //    .getItem(/** mapping value needed */);

    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getAppearance() {
        return appearance;
    }

    public void setAppearance(int appearance) {
        this.appearance = appearance;
    }

    public int getEnergy() {
        return energy;
    }

    public void setEnergy(int energy) {
        this.energy = energy;
    }

    public Address getAddress() {
        return address;
    }

    public void setAddress(Address address) {
        this.address = address;
    }

    public void eat() {}

    public void sleep() {
        System.out.println("sleep");
    }

    public void speak() {
        System.out.println("speak");
    }

    public void feed() {}
}
