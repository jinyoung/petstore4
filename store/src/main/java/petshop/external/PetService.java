package petshop.external;

import java.util.Date;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

@FeignClient(name = "pet", url = "http://pet:8080")
public interface PetService {
    @RequestMapping(method = RequestMethod.GET, path = "/pets/{id}")
    public Pet getPet(@PathVariable("id") Long id);
}
