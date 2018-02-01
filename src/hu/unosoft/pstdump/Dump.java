package pstdump;

import java.io.OutputStreamWriter;
import java.io.InputStream;
import java.io.OutputStream;
import java.io.ByteArrayOutputStream;

import java.util.Vector;
import java.util.Base64;

import com.pff.PSTException;
import com.pff.PSTFile;
import com.pff.PSTFolder;
import com.pff.PSTMessage;
import com.pff.PSTRecipient;
import com.pff.PSTAttachment;

/*import javax.mail.Header;
import javax.mail.MessagingException;
import javax.mail.Session;
import javax.mail.internet.InternetAddress;
import javax.mail.internet.InternetHeaders;
import javax.mail.internet.MimeBodyPart;
import javax.mail.internet.MimeMessage;
import javax.mail.internet.MimeMultipart;
import javax.mail.internet.MimeMessage.RecipientType;
import javax.mail.util.ByteArrayDataSource;
*/

//import com.cedarsoftware.util.io.JsonWriter;
import com.google.gson.stream.JsonWriter;

public class Dump {
	OutputStreamWriter out = null;
    public static void main(final String[] args) {
        new Dump(args[0]);
    }

    public Dump(final String filename) {
        try {
			out = new OutputStreamWriter(System.out, "UTF-8");
            final PSTFile pstFile = new PSTFile(filename);
            this.processFolder(pstFile.getRootFolder());
        } catch (final Exception err) {
            err.printStackTrace();
        }
    }

    public void processFolder(final PSTFolder folder) throws PSTException, java.io.IOException {
        // go through the folders...
        if (folder.hasSubfolders()) {
            final Vector<PSTFolder> childFolders = folder.getSubFolders();
            for (final PSTFolder childFolder : childFolders) {
                this.processFolder(childFolder);
            }
        }

        // and now the emails for this folder
        if (folder.getContentCount() > 0) {
            PSTMessage email = (PSTMessage) folder.getNextChild();
            while (email != null) {
				writeMessage(folder.getDisplayName(), email, out);

                email = (PSTMessage) folder.getNextChild();
            }
        }
    }

	public void writeMessage(String folder, PSTMessage msg, java.io.OutputStreamWriter out) throws java.io.IOException, PSTException  {
		JsonWriter jw = new JsonWriter(out);
		jw.beginObject();
		jw.name("Folder").value(folder);
		jw.name("Body").value(msg.getBody());
		jw.name("BodyHTML").value(msg.getBodyHTML());
		jw.name("BodyPrefix").value(msg.getBodyPrefix());
		jw.name("ClientSubmitTime").value(msg.getClientSubmitTime().toString());
		byte[] b = msg.getConversationId();
		if(b != null && b.length != 0) {
			jw.name("ConversationID").value(Base64.getEncoder().encodeToString(b));
		}
		jw.name("ConversationTopic").value(msg.getConversationTopic());
		jw.name("InReplyTo").value(msg.getInReplyToId());
		jw.name("ArticleNumber").value(msg.getInternetArticleNumber());
		jw.name("ID").value(msg.getInternetMessageId());
		jw.name("Class").value(msg.getMessageClass());
		jw.name("DeliveryTime").value(msg.getMessageDeliveryTime().toString());
		jw.name("Size").value(msg.getMessageSize());
		jw.name("NextSendAcct").value(msg.getNextSendAcct());
		jw.name("PrimarySendAccount").value(msg.getPrimarySendAccount());
		jw.name("ReceivedRepresenting").beginObject();
		jw.name("Address").value(msg.getRcvdRepresentingEmailAddress());
		jw.name("Name").value(msg.getRcvdRepresentingName());
		jw.endObject();
		jw.name("ReceivedBy").beginObject();
		jw.name("Address").value(msg.getReceivedByAddress());
		jw.name("Name").value(msg.getReceivedByName());
		jw.endObject();
		jw.name("ReturnPath").value(msg.getReturnPath());
		jw.name("RTFBody").value(msg.getRTFBody());
		jw.name("Sender").beginObject();
		jw.name("Address").value(msg.getSenderEmailAddress());
		jw.name("Name").value(msg.getSenderName());
		jw.endObject();
		jw.name("SentRepresenting").beginObject();
		jw.name("Address").value(msg.getSentRepresentingEmailAddress());
		jw.name("Name").value(msg.getSentRepresentingName());
		jw.endObject();
		jw.name("Subject").value(msg.getSubject());
		jw.name("Headers").value(msg.getTransportMessageHeaders());
		jw.name("CompName").value(msg.getURLCompName());
		jw.name("BCC").value(msg.getDisplayBCC());
		jw.name("CC").value(msg.getDisplayCC());
		jw.name("To").value(msg.getDisplayTo());

		jw.name("Recipients").beginArray();
		for(int i=0; i < msg.getNumberOfRecipients(); i++) {
			jw.beginObject();
			PSTRecipient rcpt = msg.getRecipient(i);
			jw.name("Address").value(rcpt.getEmailAddress());
			jw.name("Name").value(rcpt.getDisplayName());
			jw.endObject();
		}
		jw.endArray();

		jw.name("Attachments").beginArray();
		for(int i=0; i < msg.getNumberOfAttachments(); i++) {
			jw.beginObject();
			PSTAttachment att = msg.getAttachment(i);
			jw.name("ContentDisposition").value(att.getAttachmentContentDisposition());
			jw.name("Size").value(att.getAttachSize());
			jw.name("ID").value(att.getContentId());
			jw.name("Created").value(att.getCreationTime().toString());
			jw.name("FileName").value(att.getFilename());
			jw.name("FileSize").value(att.getFilesize());
			jw.name("ContentType").value(att.getMimeTag());
			jw.name("Data").value(toB64String(att.getFileInputStream()));
			jw.endObject();
		}
		jw.endArray();

		jw.endObject();
		out.write('\n');
	}

	public String toB64String(InputStream is) throws java.io.IOException {
		ByteArrayOutputStream buffer = new ByteArrayOutputStream();
		OutputStream out = Base64.getEncoder().wrap(buffer);
		int nRead;
		byte[] data = new byte[16384];

		while ((nRead = is.read(data, 0, data.length)) != -1) {
			  out.write(data, 0, nRead);
		}

		out.close();
		buffer.flush();
		return buffer.toString();
	}

}

