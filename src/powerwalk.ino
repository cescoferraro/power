#include <SPI.h>
#include "Wire.h"


#define SHIELD_1_I2C_ADDRESS  0x20  // 0x20 is the address with all jumpers removed
#define SHIELD_2_I2C_ADDRESS  0x21  // 0x21 is the address with a jumper on position A0
#define SHIELD_3_I2C_ADDRESS  0x22  // 0x21 is the address with a jumper on position A0
#define SHIELD_4_I2C_ADDRESS  0x23  // 0x21 is the address with a jumper on position A0
#define SHIELD_5_I2C_ADDRESS  0x24  // 0x21 is the address with a jumper on position A0
#define SHIELD_6_I2C_ADDRESS  0x25  // 0x21 is the address with a jumper on position A0
#define SHIELD_7_I2C_ADDRESS  0x26  // 0x21 is the address with a jumper on position A0
#define SHIELD_8_I2C_ADDRESS  0x27  // 0x21 is the address with a jumper on position A0
#define MAC_I2C_ADDRESS       0x50  // Microchip 24AA125E48 I2C ROM address
byte shield1BankA = 0; // Current status of all outputs on first shield, one bit per output
byte shield2BankA = 0; // Current status of all outputs on second shield, one bit per output
byte shield3BankA = 0; // Current status of all outputs on second shield, one bit per output
byte shield4BankA = 0; // Current status of all outputs on second shield, one bit per output
byte shield5BankA = 0; // Current status of all outputs on second shield, one bit per output
byte shield6BankA = 0; // Current status of all outputs on second shield, one bit per output
byte shield7BankA = 0; // Current status of all outputs on second shield, one bit per output
byte shield8BankA = 0; // Current status of all outputs on second shield, one bit per output



const unsigned long sampleTime = 100000UL;
const unsigned long numSamples = 250UL;
const unsigned long sampleInterval = sampleTime / numSamples;
const int adc_zero = 514;
int ac712;



String cmd;
boolean cmdRec = false;
int debugg = 0;


void setup() {




  //  PROTOSHIELD
  Wire.begin(); // Wake up I2C bus
  Serial.begin(9600);


  //RELAY 8 I2C CONNECTIONS
  initialiseShield(SHIELD_1_I2C_ADDRESS);
  sendRawValueToLatch1(0);
  initialiseShield(SHIELD_2_I2C_ADDRESS);
  sendRawValueToLatch2(0);
  initialiseShield(SHIELD_3_I2C_ADDRESS);
  sendRawValueToLatch3(0);



  initialiseShield(SHIELD_4_I2C_ADDRESS);
  sendRawValueToLatch4(0);
  initialiseShield(SHIELD_5_I2C_ADDRESS);
  sendRawValueToLatch5(0);
  initialiseShield(SHIELD_6_I2C_ADDRESS);
  sendRawValueToLatch6(0);
      initialiseShield(SHIELD_7_I2C_ADDRESS);
    sendRawValueToLatch7(0);
  //    initialiseShield(SHIELD_8_I2C_ADDRESS);
  //  sendRawValueToLatch8(0);


  //  LOOP DEFINITION

  startlights();



}


void loop() {
  handleCmd();

}


void serialEvent() {

  while (Serial.available() > 0) {

    char inByte = (char)Serial.read();
    if (inByte == ':') {
      cmdRec = true;
      return;
    }
    else if (inByte == '@') {
      cmd = "";
      cmdRec = false;
      return;
    }
    else {
      cmd += inByte;
      return;
    }
  }
}







void handleCmd() {
  if (!cmdRec) return;

  // If you have problems try changing this value,
  // my MEGA2560 has a lot of space
  int data[15];
  int numArgs = 0;

  int beginIdx = 0;
  int idx = cmd.indexOf(",");

  String arg;
  char charBuffer[20];


  while (idx != -1) {
    arg = cmd.substring(beginIdx, idx);
    arg.toCharArray(charBuffer, 16);

    data[numArgs++] = atoi(charBuffer);
    beginIdx = idx + 1;
    idx = cmd.indexOf(",", beginIdx);
  }
  // And also fetch the last command
  arg = cmd.substring(beginIdx);
  arg.toCharArray(charBuffer, 16);
  data[numArgs++] = atoi(charBuffer);
  // Now execute the command
  execCmd(data);

  cmdRec = false;
}



void execCmd(int* data) {

  if ( debugg == 1 )
  {
    Serial.print("mode: ");
    Serial.print(data[0]);
    Serial.print("port or value: ");
    Serial.print(data[1]);
    Serial.print("value: ");
    Serial.println(data[2]);
  }

  switch (data[0]) {



    case 101:
      {
        for (int i = 2; i < (data[1] * 2) + 1; i += 2) {
          pinMode(data[i], OUTPUT);
          analogWrite(data[i], data[i + 1]);
        }
      }
      break;

    case 102:
      {

        pinMode(data[1], INPUT);
        int sensor = analogRead(data[1]);
        Serial.print(sensor);

        //for (int i=0; i <= 10; i++){
        //Serial.print("@");
        //
        //}
        //Serial.print("@");




      }
      break;

    case 103:
      {
        String result = "";
        int sensor = 0;
        for (int j = 2; j < data[1] + 2; j++) {
          pinMode(data[j], INPUT);
          sensor = analogRead(data[j]);
          result += String(sensor) + ",";
        }
        Serial.println(result);
      }
      break;

    ////// LIGA/DESLIGA CANAL 8 @104,8,255:  @104,8,0:

    case 104:
      {
        if ( data[2] == 255 ) {
          setLatchChannelOn(data[1]);
        };
        if ( data[2] == 0 ) {
          setLatchChannelOff(data[1]);
        };



      }
      break;

    case 105:
      {
        int bruna = data[1];
        //        readsensor(bruna);
      }
      break;

    case 106:
      {
        int port = data[1];
        if (port == 1) {
          readPortsI();
        }
        if (port == 2) {
          readPortsII();
        }
        if (port == 3) {
          readPortsIII();
        }
        if (port == 4) {
          readPortsIV();
        }
        if (port == 5) {
          readPortsV();
        }
        if (port == 6) {
          readPortsVI();
        }
              if (port == 7) {
          readPortsVII();
        }
                      if (port == 8) {
          readPortsVIII();
        }
      }
      break;


    default:
      {
        pinMode(data[0], OUTPUT);
        analogWrite(data[0], data[1]);
        if ( debugg == 1 )
        {
          Serial.print("PIN | ");
          Serial.println(data[0]);
        }
      }
      break;
  }
}

byte readRegister(byte r)
{
  unsigned char v;
  Wire.beginTransmission(MAC_I2C_ADDRESS);
  Wire.write(r);  // Register to read
  Wire.endTransmission();

  Wire.requestFrom(MAC_I2C_ADDRESS, 1); // Read a byte
  while (!Wire.available())
  {
    // Wait
  }
  v = Wire.read();
  return v;
}


void initialiseShield(int shieldAddress)
{
  // Set addressing style
  Wire.beginTransmission(shieldAddress);
  Wire.write(0x12);
  Wire.write(0x20); // use table 1.4 addressing
  Wire.endTransmission();

  // Set I/O bank A to outputs
  Wire.beginTransmission(shieldAddress);
  Wire.write(0x00); // IODIRA register
  Wire.write(0x00); // Set all of bank A to outputs
  Wire.endTransmission();
}






void setLatchChannelOn(byte channelId)
{


  //  Serial.println("channelId");
  //  Serial.println(channelId);

  if ( channelId >= 1 && channelId <= 8 )
  {
    byte shieldOutput = channelId;
    byte channelMask = 1 << (shieldOutput - 1);
    shield1BankA = shield1BankA | channelMask;
    sendRawValueToLatch1(shield1BankA);
    //      Serial.println("shieldOutput");
    //  Serial.println(shieldOutput);
    //  Serial.println("channelMask");
    //  Serial.println(channelMask);
  }
  else if ( channelId >= 9 && channelId <= 16 )
  {

    byte shieldOutput = channelId - 8;

    byte channelMask = 1 << (shieldOutput - 1);
    shield2BankA = shield2BankA | channelMask;
    sendRawValueToLatch2(shield2BankA);

  }
  else if ( channelId >= 17 && channelId <= 24 )
  {

    byte shieldOutput = channelId - 16;

    byte channelMask = 1 << (shieldOutput - 1);

    shield3BankA = shield3BankA | channelMask;
    sendRawValueToLatch3(shield3BankA);

  }
  else if ( channelId >= 25 && channelId <= 32 )
  {

    byte shieldOutput = channelId - 24;

    byte channelMask = 1 << (shieldOutput - 1);

    shield4BankA = shield4BankA | channelMask;
    sendRawValueToLatch4(shield4BankA);

  }
  else if ( channelId >= 33 && channelId <= 40 )
  {

    byte shieldOutput = channelId - 32;

    byte channelMask = 1 << (shieldOutput - 1);

    shield5BankA = shield5BankA | channelMask;
    sendRawValueToLatch5(shield5BankA);

  }
  else if ( channelId >= 41 && channelId <= 48 )
  {

    byte shieldOutput = channelId - 40;

    byte channelMask = 1 << (shieldOutput - 1);

    shield6BankA = shield6BankA | channelMask;
    sendRawValueToLatch6(shield6BankA);

  }
    else if ( channelId >= 49 && channelId <= 56 )
  {
    byte shieldOutput = channelId - 48;
    byte channelMask = 1 << (shieldOutput - 1);
    shield7BankA = shield7BankA | channelMask;
    sendRawValueToLatch7(shield7BankA);

  }
    else if ( channelId >= 57 && channelId <= 64 )
  {
    byte shieldOutput = channelId - 56;
    byte channelMask = 1 << (shieldOutput - 1);
    shield8BankA = shield8BankA | channelMask;
    sendRawValueToLatch8(shield8BankA);

  }




}


void setLatchChannelOff (byte channelId)
{
  if ( channelId >= 1 && channelId <= 8 )
  {
    byte shieldOutput = channelId;
    byte channelMask = 255 - ( 1 << (shieldOutput - 1));
    shield1BankA = shield1BankA & channelMask;
    sendRawValueToLatch1(shield1BankA);
  }
  else if ( channelId >= 9 && channelId <= 16 )
  {
    byte shieldOutput = channelId - 8;
    byte channelMask = 255 - ( 1 << (shieldOutput - 1));
    shield2BankA = shield2BankA & channelMask;
    sendRawValueToLatch2(shield2BankA);
  }
  else if ( channelId >= 17 && channelId <= 24 )
  {
    byte shieldOutput = channelId - 16;
    byte channelMask = 255 - ( 1 << (shieldOutput - 1));
    shield3BankA = shield3BankA & channelMask;
    sendRawValueToLatch3(shield3BankA);
  }
  else if ( channelId >= 25 && channelId <= 32 )
  {
    byte shieldOutput = channelId - 24;
    byte channelMask = 255 - ( 1 << (shieldOutput - 1));
    shield4BankA = shield4BankA & channelMask;
    sendRawValueToLatch4(shield4BankA);
  }
  else if ( channelId >= 33 && channelId <= 40 )
  {
    byte shieldOutput = channelId - 32;
    byte channelMask = 255 - ( 1 << (shieldOutput - 1));
    shield5BankA = shield5BankA & channelMask;
    sendRawValueToLatch5(shield5BankA);
  }
  else if ( channelId >= 41 && channelId <= 48 )
  {
    byte shieldOutput = channelId - 40;
    byte channelMask = 255 - ( 1 << (shieldOutput - 1));
    shield6BankA = shield6BankA & channelMask;
    sendRawValueToLatch6(shield6BankA);
  }
    else if ( channelId >= 49 && channelId <= 56 )
  {
    byte shieldOutput = channelId - 48;
    byte channelMask = 255 - ( 1 << (shieldOutput - 1));
    shield7BankA = shield7BankA & channelMask;
    sendRawValueToLatch7(shield7BankA);
  }
      else if ( channelId >= 57 && channelId <= 64 )
  {
    byte shieldOutput = channelId - 56;
    byte channelMask = 255 - ( 1 << (shieldOutput - 1));
    shield8BankA = shield8BankA & channelMask;
    sendRawValueToLatch8(shield8BankA);
  }
}


void sendRawValueToLatch1(byte rawValue)
{
  Wire.beginTransmission(SHIELD_1_I2C_ADDRESS);
  Wire.write(0x12);        // Select GPIOA
  Wire.write(rawValue);    // Send value to bank A
  shield1BankA = rawValue;
  Wire.endTransmission();
}

void sendRawValueToLatch2(byte rawValue)
{
  Wire.beginTransmission(SHIELD_2_I2C_ADDRESS);
  Wire.write(0x12);        // Select GPIOA
  Wire.write(rawValue);    // Send value to bank A
  shield2BankA = rawValue;
  Wire.endTransmission();
}


void sendRawValueToLatch3(byte rawValue)
{
  Wire.beginTransmission(SHIELD_3_I2C_ADDRESS);
  Wire.write(0x12);        // Select GPIOA
  Wire.write(rawValue);    // Send value to bank A
  shield3BankA = rawValue;
  Wire.endTransmission();
}


void sendRawValueToLatch4(byte rawValue)
{
  Wire.beginTransmission(SHIELD_4_I2C_ADDRESS);
  Wire.write(0x12);        // Select GPIOA
  Wire.write(rawValue);    // Send value to bank A
  shield4BankA = rawValue;
  Wire.endTransmission();
}
void sendRawValueToLatch5(byte rawValue)
{
  Wire.beginTransmission(SHIELD_5_I2C_ADDRESS);
  Wire.write(0x12);        // Select GPIOA
  Wire.write(rawValue);    // Send value to bank A
  shield5BankA = rawValue;
  Wire.endTransmission();
}
void sendRawValueToLatch6(byte rawValue)
{
  Wire.beginTransmission(SHIELD_6_I2C_ADDRESS);
  Wire.write(0x12);        // Select GPIOA
  Wire.write(rawValue);    // Send value to bank A
  shield6BankA = rawValue;
  Wire.endTransmission();
}
void sendRawValueToLatch7(byte rawValue)
{
  Wire.beginTransmission(SHIELD_7_I2C_ADDRESS);
  Wire.write(0x12);        // Select GPIOA
  Wire.write(rawValue);    // Send value to bank A
  shield7BankA = rawValue;
  Wire.endTransmission();
}
void sendRawValueToLatch8(byte rawValue)
{
  Wire.beginTransmission(SHIELD_8_I2C_ADDRESS);
  Wire.write(0x12);        // Select GPIOA
  Wire.write(rawValue);    // Send value to bank A
  shield8BankA = rawValue;
  Wire.endTransmission();
}

void readPortsI()
{
  byte inputs = 0;

  Wire.beginTransmission(SHIELD_1_I2C_ADDRESS);
  Wire.write(0x12); // set MCP23017 memory pointer to GPIOA address
  Wire.endTransmission();
  Wire.requestFrom(0x20, 1); // request one byte of data from MCP20317
  inputs = Wire.read(); // store the incoming byte into "inputs"

  Serial.println(inputs, BIN); // display the contents of the GPIOA register in binary
}

void readPortsII()
{
  byte inputs = 0;

  Wire.beginTransmission(SHIELD_2_I2C_ADDRESS);
  Wire.write(0x12); // set MCP23017 memory pointer to GPIOA address
  Wire.endTransmission();
  Wire.requestFrom(0x21, 1); // request one byte of data from MCP20317
  inputs = Wire.read(); // store the incoming byte into "inputs"

  Serial.println(inputs, BIN); // display the contents of the GPIOA register in binary
}
void readPortsIII()
{
  byte inputs = 0;

  Wire.beginTransmission(SHIELD_3_I2C_ADDRESS);
  Wire.write(0x12); // set MCP23017 memory pointer to GPIOA address
  Wire.endTransmission();
  Wire.requestFrom(0x22, 1); // request one byte of data from MCP20317
  inputs = Wire.read(); // store the incoming byte into "inputs"

  Serial.println(inputs, BIN); // display the contents of the GPIOA register in binary
}

void readPortsIV()
{
  byte inputs = 0;

  Wire.beginTransmission(SHIELD_4_I2C_ADDRESS);
  Wire.write(0x12); // set MCP23017 memory pointer to GPIOA address
  Wire.endTransmission();
  Wire.requestFrom(0x23, 1); // request one byte of data from MCP20317
  inputs = Wire.read(); // store the incoming byte into "inputs"

  Serial.println(inputs, BIN); // display the contents of the GPIOA register in binary
}

void readPortsV()
{
  byte inputs = 0;

  Wire.beginTransmission(SHIELD_5_I2C_ADDRESS);
  Wire.write(0x12); // set MCP23017 memory pointer to GPIOA address
  Wire.endTransmission();
  Wire.requestFrom(0x24, 1); // request one byte of data from MCP20317
  inputs = Wire.read(); // store the incoming byte into "inputs"

  Serial.println(inputs, BIN); // display the contents of the GPIOA register in binary
}

void readPortsVI()
{
  byte inputs = 0;

  Wire.beginTransmission(SHIELD_6_I2C_ADDRESS);
  Wire.write(0x12); // set MCP23017 memory pointer to GPIOA address
  Wire.endTransmission();
  Wire.requestFrom(0x25, 1); // request one byte of data from MCP20317
  inputs = Wire.read(); // store the incoming byte into "inputs"

  Serial.println(inputs, BIN); // display the contents of the GPIOA register in binary
}
void readPortsVII()
{
  byte inputs = 0;

  Wire.beginTransmission(SHIELD_7_I2C_ADDRESS);
  Wire.write(0x12); // set MCP23017 memory pointer to GPIOA address
  Wire.endTransmission();
  Wire.requestFrom(0x26, 1); // request one byte of data from MCP20317
  inputs = Wire.read(); // store the incoming byte into "inputs"

  Serial.println(inputs, BIN); // display the contents of the GPIOA register in binary
}
void readPortsVIII()
{
  byte inputs = 0;

  Wire.beginTransmission(SHIELD_8_I2C_ADDRESS);
  Wire.write(0x12); // set MCP23017 memory pointer to GPIOA address
  Wire.endTransmission();
  Wire.requestFrom(0x27, 1); // request one byte of data from MCP20317
  inputs = Wire.read(); // store the incoming byte into "inputs"

  Serial.println(inputs, BIN); // display the contents of the GPIOA register in binary
}






void startlights() {

  for (int i = 1; i <= 48; i++) {
    setLatchChannelOn(i);
  }


}
void stoplights() {
  for (int i = 1; i <= 48; i++) {
    setLatchChannelOff(i);
  }
}