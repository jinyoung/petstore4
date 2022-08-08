package fooddelivery.domain;

import fooddelivery.domain.*;
import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

@RepositoryRestResource(collectionResourceRel = "details", path = "details")
public interface DetailRepository
    extends PagingAndSortingRepository<Detail, Long> {}
