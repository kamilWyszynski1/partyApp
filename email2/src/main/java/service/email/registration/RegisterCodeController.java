package service.email.registration;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.*;
import org.springframework.http.converter.json.MappingJackson2HttpMessageConverter;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.HttpClientErrorException;
import org.springframework.web.client.RestTemplate;

import javax.mail.MessagingException;
import java.util.List;
import java.util.Random;

@CrossOrigin(origins = "http://localhost:8080")
@RestController
public class RegisterCodeController {

    @Value("${access_token}")
    private String access_token;

    private final RegisterCodeRepository registerCodeRepository;

    private final EmailServiceImpl emailService;

    public RegisterCodeController(RegisterCodeRepository registerCodeRepository, EmailServiceImpl emailService) {
        this.registerCodeRepository = registerCodeRepository;
        this.emailService = emailService;
    }

    @RequestMapping(value = "/verification", method = RequestMethod.POST)
    public RegisterCode createRegisterCode(@RequestBody RegisterCode registerCode) {
        int code = new Random().nextInt(900000) + 100000;
        registerCode.setCode(code);

        try {
            emailService.sendMessage("kamil-wyszynski@wp.pl", String.valueOf(code));
        } catch (MessagingException e) {
            e.printStackTrace();
        }

        registerCodeRepository.save(registerCode);
        return registerCode;
    }

    @RequestMapping(value = "/verification", method = RequestMethod.GET)
    public List<RegisterCode> getCodes(){
        return (List<RegisterCode>) registerCodeRepository.findAll();
    }

    @RequestMapping(value = "/verification", method = RequestMethod.DELETE)
    public void deleteCodes(){
        registerCodeRepository.deleteAll();
    }

    @RequestMapping(value = "/verify", method = RequestMethod.POST)
    public ResponseEntity<RegisterCode> verify(@RequestBody RegisterCode registerCode){
        /*
            Endpoint verify if sent code matches email, if so it sends request to activate
            user to User Service and return STATUS 200. If Code doesnt match email it returns
            404 or 401.
         */

        RegisterCode registerCodeDb = registerCodeRepository.findById(registerCode.getEmail())
                .orElse(null);

        if (registerCodeDb == null)
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);

        if (registerCodeDb.getCode() == registerCode.getCode()) {
            // Delete verified email
            registerCodeRepository.delete(registerCode);

            try {
                this.activate_user(registerCode.getEmail());
                return new ResponseEntity<>(HttpStatus.OK);
            }
            catch(HttpClientErrorException error){
                return new ResponseEntity<>(HttpStatus.NOT_FOUND);
            }
        }
        else
            return new ResponseEntity<>(HttpStatus.UNAUTHORIZED);
    }

    @RequestMapping(value = "/activate", method = RequestMethod.GET)
    private HttpStatus activate_user(String email) throws HttpClientErrorException{
        /*
            Function sends request to User Service and call activate user function
         */
        final String url = "http://localhost:8000/activate/";

        RestTemplate restTemplate = new RestTemplate();
        restTemplate.getMessageConverters().add(new MappingJackson2HttpMessageConverter());

        HttpHeaders headers = new HttpHeaders();
        headers.set("Authorization", "Bearer tokenik");
        headers.setContentType(MediaType.APPLICATION_JSON);

        String requestJson = "{\"email\":\"test@wp.pl\"}" ;
        String.format(requestJson, email);
        HttpEntity<String> entity = new HttpEntity<>(requestJson, headers);

        ResponseEntity<String> result = restTemplate.postForEntity(url, entity, String.class);

        return result.getStatusCode();
    }
}
