package service.email.registration;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.mail.javamail.MimeMessageHelper;
import org.springframework.stereotype.Component;

import javax.mail.MessagingException;
import javax.mail.internet.MimeMessage;

@Component
public class EmailServiceImpl {

    @Qualifier("javaMailSender")
    @Autowired
    private JavaMailSender emailSender;

    private SimpleMailMessage message = new SimpleMailMessage();


    private final String template = " <body style='font-size: 22px;'> Here's your registration code: <b> %s </b>, " +
            "it will expire in 5min!</body>";

    public void sendMessage(String to, String text) throws MessagingException {
        MimeMessage mimeMessage = emailSender.createMimeMessage();
        MimeMessageHelper helper = new MimeMessageHelper(mimeMessage, false, "utf-8");
        String htmlMsg = String.format(template, text);
        mimeMessage.setContent(htmlMsg, "text/html");
        helper.setTo(to);
        helper.setSubject("Registration Code");

        emailSender.send(mimeMessage);
    }
}
