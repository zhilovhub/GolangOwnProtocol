# Спецификация протокола IlyaProtocol

**ILYAPROTOCOL PACKET STRUCTURE**  
(offset: 0) HEADER (3 bytes) [`0xDD`, `0xEF`, `0xDD`]  
(offset: 3) PACKET TYPE (1 byte)  
(offset: 4) PACKET SUBTYPE (1 byte)  
(offset: 5) FIELDS (FIELD[])  
(offset: END) PACKET ENDING (2 bytes) [`0x00`, `0xFF`]  

**FIELD STRUCTURE**  
(offset: 0) FIELD_ID (1 byte)  
(offset: 1) FIELD SIZE (1 byte)  
(offset: 2) FIELD CONTENTS
